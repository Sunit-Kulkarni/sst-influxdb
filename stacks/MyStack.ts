import { StackContext, Api, NextjsSite } from "sst/constructs";

export function API({ stack }: StackContext) {
  const api = new Api(stack, "api", {
    customDomain: "api-iot-dev.sunitkulkarni.com",
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
          INFLUXDB_TOKEN: process.env.INFLUXDB_TOKEN,
          DATABASE_URL: process.env.DATABASE_URL
        }
      }
    }
  });
  const site = new NextjsSite(stack, "countryclub", {
    customDomain: {
      domainName: "countryclub.sunitkulkarni.com",
      hostedZone: "sunitkulkarni.com"
    },
    path: "packages/web"
  })
  stack.addOutputs({
    ApiEndpoint: api.url,
    URL: site.url,
  });
}
