package pi

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
}

type AuthToken struct {
	Token  string
	Expiry time.Time
}

type mpesa struct {
	httpClient platform.HTTPClient
	config     state.MPesaConfig
	log        loggres.logger
	AuthToken
}

func Init(httpClient platform.HTTPClient, config state.MPesaConfig, log logit.Logger) platform.MPesa {

	mpesa := mpesa{
		httpClient: httpClient,
		config:     config,
		log:        log,
	}

	// if err := mpesa.refresh(context.Background()); err != nil {
	// 	log.Fatal(context.Background(), fmt.Sprintf("Unable to generate a token: %v", err))
	// }

	return &mpesa
}
func (m *mpesa) InitiateOC(ctx context.Context, param dto.MpesaOcInitiate,
	paymentID string) (*dto.MpesaOcInitiateResponse, error) {
	if m.isExpired() {
		if err := m.refresh(ctx); err != nil {
			return nil, err
		}
	}

	if err := param.Validate(); err != nil {
		err := errors.ErrMpesaConn.Wrap(err, "required fields missing")
		m.log.Error(ctx, "required fields missing", zap.Error(err))
		return nil, err
	}

	param.BusinessShortCode = m.config.BusinessShortCode
	param.PartyB = m.config.BusinessShortCode
	param.Password = m.config.Passkey
	param.Timestamp = utils.GenerateTimestamp(time.Now())
	param.TransactionType = m.config.TransactionType
	param.CallBackURL = fmt.Sprintf(m.config.OnlineCheckoutCallbackURL,
		paymentID)

	response := dto.MpesaOcInitiateResponse{}

	_, err := m.httpClient.DoRequest(
		ctx, http.MethodPost,
		m.config.OnlineCheckoutInitiateURL,
		"application/json",
		func(r *http.Request) {
			r.Header.Set("Authorization", "Bearer "+m.AuthToken.Token)
		},
		param,
		&response,
	)
	if err != nil {
		err := errors.ErrMpesaConn.Wrap(err, "unable to initiate the payment process")
		m.log.Error(ctx, "unable to initiate the payment process", zap.Error(err))
		return nil, err
	}
	if response.ResponseCode != "0" {
		err := errors.ErrMpesaConn.New(response.ResponseDescription)
		m.log.Error(ctx,
			"error initiating the payment process",
			zap.Error(err), zap.String("response",
				response.ResponseDescription))
		return nil, err
	}

	return &response, nil
}
