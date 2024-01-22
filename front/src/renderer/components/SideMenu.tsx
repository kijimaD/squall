import { Container, Grid, Paper, Button } from "@mui/material";
import { HeaderLogo } from "./HeaderLogo";
import { EntryButton } from "./EntryButton";
import { Counter } from "./Counter";
import AddCircleOutlineIcon from "@mui/icons-material/AddCircleOutline";
import { useDispatch, useSelector } from "react-redux";
import { View, add } from "../redux/viewSlice";

export const SideMenu = () => {
  const views = useSelector((state) => state.view.views);
  const dispatch = useDispatch();

  const newView = async () => {
    try {
      const id = await window.myAPI.invoke("openNewView", {});
      dispatch(add([id]));
    } catch (error) {
      console.error("error:", error.message);
    }
  };

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
            <EntryButton title="Home" url="main_window" />
            <EntryButton title="Google" url="google.com" />
            <EntryButton title="Amazon" url="amazon.com" />
            {views.map((v, i) => {
              return <p>{v.viewId}</p>;
            })}
          </Grid>
        </Grid>
      </Container>
      <Counter />
    </Container>
  );
};
