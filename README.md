# Notes SLS

## Prerequisites

Before getting started with this serverless application, make sure you have the following prerequisites:

- [AWS account](https://aws.amazon.com/) for deployment
- [Serverless Framework](https://www.serverless.com/framework/docs/getting-started)
- [Go](https://go.dev/)

## Configuration

Before deploying the application, you need to configure your serverless framework with your AWS credentials and other environment-specific settings. Make sure to update the following configuration files:

**Local enviroment variables**:
```
AWS_ACCESS_KEY_ID={{your_access_key_id}}
AWS_SECRET_ACCESS_KEY={{your_secret_access_key}}
```

## Deployment

Build and deploy the entire serverless application
```
make deploy
```

Clean up generated files and remove all deployed resources of the serverless application
```
make remove
```

Run `make help` to discover more useful commands 
