import { StackContext, Api, NextjsSite } from "sst/constructs";

export function API({ stack }: StackContext) {
  const api = new Api(stack, "api", {
    customDomain: {
      domainName: process.env.API_DOMAIN_NAME,
      hostedZone: process.env.HOSTED_ZONE,
    },
    routes: {
      "GET /": "packages/functions/src/main.go",
      "POST /sensor": "packages/functions/lambdas/create/create.go",
      "POST /sensor_batch": "packages/functions/lambdas/batch/batch.go",
      "POST /member": "packages/functions/lambdas/member/create/create.go",
      "GET /member/{id}": "packages/functions/lambdas/member/get/get.go",
    },
    defaults: {
      function: {
        environment: {
          // @ts-ignore
          INFLUXDB_TOKEN: process.env.INFLUXDB_TOKEN,
          // @ts-ignore
          DATABASE_URL: process.env.DATABASE_URL
        }
      }
    }
  });

  const site = new NextjsSite(stack, "countryclub", {
    customDomain: {
      domainName: process.env.COUNTRY_CLUB_DOMAIN_NAME,
      hostedZone: process.env.HOSTED_ZONE,
    },
    path: "packages/web"
  })

  stack.addOutputs({
    ApiEndpoint: api.url,
    URL: site.url,
  });
}
