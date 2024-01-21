import { Container, Grid, Paper, Button } from "@mui/material";
import { HeaderLogo } from "./HeaderLogo";
import { EntryButton } from "./EntryButton";

export const SideMenu = () => {
  return (
    <Container>
      <HeaderLogo />
      <Container maxWidth="lg" sx={{ mt: 2, mb: 2 }} className="container">
        <Grid container direction="row" spacing={2}>
          <Grid item xs={12} sm={6} spacing={1}>
            <EntryButton title="Home" url="main_window" />
            <EntryButton title="Google" url="google.com" />
            <EntryButton title="Amazon" url="amazon.com" />
          </Grid>
        </Grid>
      </Container>
    </Container>
  );
};
