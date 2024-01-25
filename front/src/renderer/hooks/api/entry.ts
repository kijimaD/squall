import {
  InfiniteData,
  QueryKey,
  UseQueryOptions,
  useInfiniteQuery,
  useMutation,
  useQuery,
  useQueryClient,
} from "@tanstack/react-query";
import { EntryApi } from "../../generated";
import { createApiConfiguration } from "./config";

const entryApi = new EntryApi(createApiConfiguration());

export const useEntries = () => {
  return useQuery({
    queryFn: async () => await entryApi.getEntries(),
    queryKey: ["entries"],
  });
};
