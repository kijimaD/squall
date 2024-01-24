import { Container, Grid, Paper, Button } from "@mui/material";
import { HeaderLogo } from "./HeaderLogo";
import { EntryButton } from "./EntryButton";
import { Counter } from "./Counter";
import AddCircleOutlineIcon from "@mui/icons-material/AddCircleOutline";
import { useDispatch, useSelector } from "react-redux";
import { View, add } from "../redux/viewSlice";
import { useEffect } from "react";

export const SideMenu = () => {
  const views = useSelector((state) => state.view.views);
  const dispatch = useDispatch();

  const newView = async () => {
    try {
      const id = await window.myAPI.invoke("openNewView", {});
      const title = await window.myAPI.invoke("getTitleById", { id });
      dispatch(add([id, title]));
    } catch (error) {
      console.error("error:", error.message);
    }
  };

  useEffect(() => {
    console.log("useEffectが実行された(on登録)");
    // ipcRendererからメッセージを受信
    window.myAPI.on("pageLoaded", (arg) => {
      console.log("メッセージ受信:", arg);
    });
  }, []);

  return (
    <Container>
      <HeaderLogo />
      <Container maxWidth="lg" sx={{ mt: 2, mb: 2 }} className="container">
        <Grid container direction="row" spacing={2}>
          <Grid item xs={12} sm={6} spacing={1}>
            <Button
              color="black"
              style={{ justifyContent: "flex-start" }}
              onClick={() => newView()}
            >
              <AddCircleOutlineIcon />
              new
            </Button>
            <Button
              color="black"
              style={{ justifyContent: "flex-start" }}
              onClick={() => {}}
            >
              <AddCircleOutlineIcon />
              load
            </Button>
            <EntryButton title="Home" url="main_window" />
            <EntryButton title="Google" url="google.com" />
            <EntryButton title="Amazon" url="amazon.com" />
            {views.map((v, i) => {
              return (
                <p>
                  {v.viewId} {v.title}
                </p>
              );
            })}
          </Grid>
        </Grid>
      </Container>
      <Counter />
    </Container>
  );
};
