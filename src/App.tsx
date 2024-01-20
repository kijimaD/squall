import CssBaseline from '@mui/material/CssBaseline'
import { Container, Grid, Paper, Button, Typography } from "@mui/material";
import { ThemeProvider, createTheme } from '@mui/material/styles'
import TornadoIcon from '@mui/icons-material/Tornado';
import "./App.css";

const theme = createTheme({
    typography: {
        button: {
            textTransform: "none"
        },
    },
    palette: {
        black: {
            main: '#333333',
        },
    },
})

export const App = () => {
    return (
        <ThemeProvider theme={theme}>
            <CssBaseline />
            <Typography variant="h3"><TornadoIcon sx={{ color: 'blue', fontSize: 34 }} />Squall</Typography>
            <Container maxWidth="lg" sx={{ mt: 2, mb: 2 }} className="container">
                <Grid container direction="row" spacing={2}>
                    <Grid item xs={12} sm={6} spacing={1}>
                        <Grid container direction="column">
                            <Button color="black"
                                    style={{justifyContent: "flex-start"}}
                                    onClick={() => {
                                        window.myAPI.invoke("changeTab", {url: "main_window"});
                                    }}>Home</Button>
                        </Grid>
                        <Grid container direction="column">
                            <Button color="black"
                                    style={{justifyContent: "flex-start"}}
                                    onClick={() => {
                                        window.myAPI.invoke("changeTab", {url: "google.com"});
                                    }}>Google</Button>
                        </Grid>
                        <Grid container direction="column">
                            <Button color="black"
                                    style={{justifyContent: "flex-start"}}
                                    onClick={() => {
                                        window.myAPI.invoke("changeTab", {url: "amazon.com"});
                                    }}>Amazon</Button>
                        </Grid>
                    </Grid>
                </Grid>
            </Container>
        </ThemeProvider>
    );
};
