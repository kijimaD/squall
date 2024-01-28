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

export const useGetEntries = (size: number, ignore_ids: number[]) => {
  return useQuery({
    queryFn: async () =>
      await entryApi.getEntries(
        size,
        ignore_ids?.length ? ignore_ids : undefined,
      ),
    queryKey: ["entries"],
  });
};
