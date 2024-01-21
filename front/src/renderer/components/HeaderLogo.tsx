import { Typography } from "@mui/material";
import TornadoIcon from "@mui/icons-material/Tornado";

export const HeaderLogo = () => {
  return (
    <Typography variant="h3">
      <TornadoIcon sx={{ color: "blue", fontSize: 34 }} />
      Squall
    </Typography>
  );
};
