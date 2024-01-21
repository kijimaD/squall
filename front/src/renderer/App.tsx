import CssBaseline from "@mui/material/CssBaseline";
import { ThemeProvider, createTheme } from "@mui/material/styles";
import { SideMenu } from "./components/SideMenu";
import "./App.css";

const theme = createTheme({
  typography: {
    button: {
      textTransform: "none",
    },
  },
  palette: {
    black: {
      main: "#333333",
    },
  },
});

export const App = () => {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <SideMenu />
    </ThemeProvider>
  );
};
