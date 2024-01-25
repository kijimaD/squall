import { Configuration } from "../../generated";

export const createApiConfiguration = () => {
  return new Configuration({
    basePath: "http://localhost:8080",
  });
};
