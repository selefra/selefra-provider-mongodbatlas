package mongodbatlas_client

import (
	"context"
	"github.com/mongodb-forks/digest"
	"go.mongodb.org/atlas/mongodbatlas"
)

func GetMongoDBAtlasClient(ctx context.Context, config *Config) (*mongodbatlas.Client, error) {
	client, err := createClient(ctx, config.PublicKey, config.PrivateKey)
	return client, err
}

func createClient(ctx context.Context, publicKey string, privateKey string) (*mongodbatlas.Client, error) {
	t := digest.NewTransport(publicKey, privateKey)
	tc, err := t.Client()
	if err != nil {
		return nil, err
	}

	return mongodbatlas.NewClient(tc), nil
	//return mongodbatlas.NewClient(&http.Client{
	//	Transport: loggingRoundTripper{tc.Transport},
	//}), nil
}
