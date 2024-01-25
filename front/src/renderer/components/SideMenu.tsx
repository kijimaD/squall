import {
  Container,
  Grid,
  Paper,
  Button,
  TextField,
  ListItemButton,
  ListItemText,
} from "@mui/material";
import { HeaderLogo } from "./HeaderLogo";
import { EntryButton } from "./EntryButton";
import HomeIcon from "@mui/icons-material/Home";
import SearchIcon from "@mui/icons-material/Search";
import CloseIcon from "@mui/icons-material/Close";
import { useDispatch, useSelector } from "react-redux";
import { View, add, update, remove } from "../redux/viewSlice";
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
            <TextField
              defaultValue={inputUrl}
              onChange={(e) => setInputUrl(e.target.value)}
              size="small"
            />
            <Button
              color="black"
              style={{ justifyContent: "flex-start" }}
              onClick={() => newView()}
            >
              <SearchIcon />
              Go
            </Button>
            <br />
            <Button
              color="black"
              style={{ justifyContent: "flex-start" }}
              onClick={() => {
                window.myAPI.invoke("changeHome", {});
              }}
            >
              <HomeIcon />
              Home
            </Button>
            {views.map((v, i) => {
              return (
                <ListItemButton>
                  <Button
                    onClick={() => {
                      window.myAPI.invoke("removeView", { id: v.viewId });
                      dispatch(remove(v.viewId));
                    }}
                  >
                    <CloseIcon />
                  </Button>
                  <ListItemText
                    primary={v.title}
                    onClick={() => {
                      window.myAPI.invoke("changeTab", { id: v.viewId });
                    }}
                  />
                </ListItemButton>
              );
            })}
          </Grid>
        </Grid>
      </Container>
    </Container>
  );
};
