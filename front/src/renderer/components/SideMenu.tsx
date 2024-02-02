import {
  Container,
  Grid,
  Paper,
  Button,
  ButtonGroup,
  TextField,
  ListItemButton,
  ListItemText,
  Tooltip,
} from "@mui/material";
import { HeaderLogo } from "./HeaderLogo";
import HomeIcon from "@mui/icons-material/Home";
import SearchIcon from "@mui/icons-material/Search";
import CloseIcon from "@mui/icons-material/Close";
import SaveIcon from "@mui/icons-material/Save";
import { useDispatch } from "react-redux";
import { View, add, updateTitle, remove } from "../redux/viewSlice";
import { myUseSelector } from "../redux/store";
import { useEffect, useState } from "react";
import { useGetEntries } from "../hooks/api/entry";
import { DoneButton } from "./DoneButton";

export const SideMenu = () => {
  const views = myUseSelector((state) => state.view.views);
  const dispatch = useDispatch();

  const newView = async () => {
    try {
      const id: number = await window.myAPI.invoke("openNewView", {
        url: inputUrl,
      });
      const v: View = { viewId: id };
      dispatch(add(v));
    } catch (error) {
      console.error("error:", error.message);
    }
  };

  useEffect(() => {
    window.myAPI.on("pageLoaded", (arg: View) => {
      const id = arg.viewId;
      const title = arg.title;
      const v: View = { viewId: id, title: title };
      dispatch(updateTitle(v));
    });
  }, []);

  const [inputUrl, setInputUrl] = useState("");
  const [currentTitle, setCurrentTitle] = useState("");
  const [reqCount, setReqCount] = useState(5);

  // TODO: レスポンスを型に入れる
  const { data, isLoading, error, refetch } = useGetEntries(
    reqCount,
    views.map((v: View) => v.dataId),
  );
  const getEntries = () => {
    data.data.map(async (v, _i) => {
      const id = await window.myAPI.invoke("openNewView", { url: v.url });
      const view: View = { viewId: id, dataId: v.id, url: v.url };
      dispatch(add(view));
    });
  };

  useEffect(() => {
    refetch();
  }, [reqCount, views]);

  return (
    <>
      <HeaderLogo />
      <Container maxWidth="lg" sx={{ mt: 2, mb: 2 }} className="container">
        <Grid container direction="row">
          <Grid item xs={12} sm={6}>
            <Container>
              <TextField value={inputUrl} size="small" />
              <Button
                style={{ justifyContent: "flex-start" }}
                onClick={() => newView()}
              >
                <SearchIcon />
              </Button>
            </Container>

            <Container>
              <TextField
                defaultValue={reqCount}
                onChange={(e) => setReqCount(Number(e.target.value))}
                size="small"
                sx={{ maxWidth: 100 }}
                type="number"
              />
              <Button onClick={(e) => getEntries()}>Load</Button>
              <Button
                style={{ justifyContent: "flex-start" }}
                onClick={() =>
                  (window.location.href =
                    "org-protocol://capture?template=L&url=" +
                    inputUrl +
                    "&title=" +
                    currentTitle)
                }
              >
                org-protocol
                <SaveIcon />
              </Button>
            </Container>

            <ListItemButton>
              <Button>
                <HomeIcon fontSize="small" />
              </Button>
              <ListItemText
                primary="Home"
                onClick={() => {
                  window.myAPI.invoke("changeHome", {});
                }}
              ></ListItemText>
            </ListItemButton>

            {views.map((v: View, i: number) => {
              return (
                <ListItemButton key={i}>
                  <ListItemText
                    primary={v.title}
                    onClick={() => {
                      window.myAPI.invoke("changeTab", { id: v.viewId });

                      const matchingView = views.find(
                        (v2: View) => v.viewId === v2.viewId,
                      );
                      if (matchingView) {
                        setInputUrl(matchingView.url);
                        setCurrentTitle(matchingView.title);
                      }
                    }}
                  />
                  <ButtonGroup
                    orientation="vertical"
                    aria-label="vertical contained button group"
                    variant="text"
                  >
                    <Tooltip title="close">
                      <Button
                        size="small"
                        onClick={() => {
                          window.myAPI.invoke("removeView", { id: v.viewId });
                          dispatch(remove(v.viewId));
                        }}
                      >
                        <CloseIcon fontSize="small" />
                      </Button>
                    </Tooltip>
                    <DoneButton viewId={v.viewId} dataId={v.dataId} />
                  </ButtonGroup>
                </ListItemButton>
              );
            })}
          </Grid>
        </Grid>
      </Container>
    </>
  );
};
