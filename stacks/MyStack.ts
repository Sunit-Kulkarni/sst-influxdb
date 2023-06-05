import { HostedZone } from "aws-cdk-lib/aws-route53";
import { Config, StackContext, Api, use, NextjsSite } from "sst/constructs";


export function API({ stack }: StackContext) {
  const INFLUXDB_TOKEN = new Config.Secret(stack, "INFLUXDB_TOKEN");
  const api = new Api(stack, "api", {
    customDomain: "api-iot-dev.sunitkulkarni.com",
    routes: {
      "GET /": "packages/functions/src/main.go",
      "POST /sensor": "packages/functions/lambdas/create/create.go",
      "POST /sensor_batch": "packages/functions/lambdas/batch/batch.go",
      "POST /member": "packages/functions/lambdas/member/create/create.go"
    },
    defaults: {
      function: {
        bind: [INFLUXDB_TOKEN],
        environment: {
          INFLUXDB_TOKEN: process.env.INFLUXDB_TOKEN
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
