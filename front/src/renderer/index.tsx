import { ThemeProvider, createTheme } from "@mui/material/styles";
import { StrictMode } from "react";
import CssBaseline from "@mui/material/CssBaseline";
import { createRoot } from "react-dom/client";
import { Provider } from "react-redux";
import { App } from "./App";
import { store } from "./redux/store";

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

createRoot(document.getElementById("app")!).render(
  <StrictMode>
    <Provider store={store}>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <App />
      </ThemeProvider>
    </Provider>
  </StrictMode>
);
