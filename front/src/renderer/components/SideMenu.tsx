import { Container, Grid, Paper, Button } from "@mui/material";
import { HeaderLogo } from "./HeaderLogo";
import { EntryButton } from "./EntryButton";
import { Counter } from "./Counter";
import AddCircleOutlineIcon from "@mui/icons-material/AddCircleOutline";
import { useDispatch, useSelector } from "react-redux";
import { View, add, update } from "../redux/viewSlice";
import { useEffect, useState } from "react";

export const SideMenu = () => {
  const views = useSelector((state) => state.view.views);
  const dispatch = useDispatch();

  const newView = async () => {
    try {
      const id = await window.myAPI.invoke("openNewView", { url: inputUrl });
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

  const [inputUrl, setInputUrl] = useState("https://github.com");

  return (
    <Container>
      <HeaderLogo />
      <Container maxWidth="lg" sx={{ mt: 2, mb: 2 }} className="container">
        <Grid container direction="row" spacing={2}>
          <Grid item xs={12} sm={6} spacing={1}>
            <input
              value={inputUrl}
              onChange={(e) => setInputUrl(e.target.value)}
            />
            <Button
              color="black"
              style={{ justifyContent: "flex-start" }}
              onClick={() => newView()}
            >
              <AddCircleOutlineIcon />
              new
            </Button>
            <br />
            <Button
              color="black"
              style={{ justifyContent: "flex-start" }}
              onClick={() => {
                window.myAPI.invoke("changeHome", {});
              }}
            >
              Home
            </Button>
            {views.map((v, i) => {
              return <EntryButton id={v.viewId} title={v.title} />;
            })}
          </Grid>
        </Grid>
      </Container>
      <Counter />
    </Container>
  );
};
