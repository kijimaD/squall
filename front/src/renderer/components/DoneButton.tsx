import { myUseSelector } from "../redux/store";
import { useDispatch } from "react-redux";
import { usePostDoneEntry } from "../hooks/api/entry";
import { Button, Tooltip } from "@mui/material";
import CheckIcon from "@mui/icons-material/Check";
import { remove } from "../redux/viewSlice";

export const DoneButton = (props: Props) => {
  const views = myUseSelector((state) => state.view.views);
  const dispatch = useDispatch();

  const { dataId } = props;
  const { viewId } = props;
  const { data, isLoading, error, refetch } = usePostDoneEntry(dataId);

  return (
    <Tooltip title="done">
      <Button
        onClick={() => {
          window.myAPI.invoke("removeView", { id: viewId });
          dispatch(remove(viewId));
          refetch();
        }}
      >
        <CheckIcon fontSize="small" color="success" />
      </Button>
    </Tooltip>
  );
};
