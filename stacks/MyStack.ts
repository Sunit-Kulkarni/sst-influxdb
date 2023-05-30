import { Config, StackContext, Api, use } from "sst/constructs";


export function API({ stack }: StackContext) {
  const INFLUXDB_TOKEN = new Config.Secret(stack, "INFLUXDB_TOKEN");
  const api = new Api(stack, "api", {
    customDomain: "api-iot-dev.sunitkulkarni.com",
    routes: {
      "GET /": "packages/functions/src/main.go",
      "POST /sensor": "packages/functions/lambdas/create.go"
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
  stack.addOutputs({
    ApiEndpoint: api.url,
  });
}
