import { StackContext, Api } from "sst/constructs";

export function API({ stack }: StackContext) {
  const api = new Api(stack, "api", {
    customDomain: "api-iot-dev.sunitkulkarni.com",
    routes: {
      "GET /": "packages/functions/src/main.go",
      "POST /sensor": "packages/functions/lambdas/create.go"
    },
  });
  stack.addOutputs({
    ApiEndpoint: api.url,
  });
}
