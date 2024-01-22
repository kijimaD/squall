import CssBaseline from "@mui/material/CssBaseline";
import { ThemeProvider, createTheme } from "@mui/material/styles";
import { SideMenu } from "./components/SideMenu";
import "./App.css";
import { store } from "./redux/store";
import { Provider } from "react-redux";

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
    <Provider store={store}>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <SideMenu />
      </ThemeProvider>
    </Provider>
  );
};
