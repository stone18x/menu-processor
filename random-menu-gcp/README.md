# Deployment

```
cd random-menu/functions
gcloud functions deploy RandomMenu --runtime go113 --trigger-http --allow-unauthenticated --project menu-processor --env-vars-file ../.env.yaml
```
