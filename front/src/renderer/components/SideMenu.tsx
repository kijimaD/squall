import { Container, Grid, Paper, Button } from "@mui/material";
import { HeaderLogo } from "./HeaderLogo";
import { EntryButton } from "./EntryButton";
import { Counter } from "./Counter";
import AddCircleOutlineIcon from "@mui/icons-material/AddCircleOutline";
import { useDispatch, useSelector } from "react-redux";
import { View, add, update } from "../redux/viewSlice";
import { useEffect } from "react";

export const SideMenu = () => {
  const views = useSelector((state) => state.view.views);
  const dispatch = useDispatch();

  const newView = async () => {
    try {
      const id = await window.myAPI.invoke("openNewView", {});
      const v: View = { viewId: id };
      dispatch(add(v));
    } catch (error) {
      console.error("error:", error.message);
    }
  };

  useEffect(() => {
    window.myAPI.on("pageLoaded", (arg) => {
      const id = arg[0];
      const title = arg[1];
      const v: View = { viewId: id, title: title };
      dispatch(update(v));
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
