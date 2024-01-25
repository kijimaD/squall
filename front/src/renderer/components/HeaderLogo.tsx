import { Typography } from "@mui/material";
import TornadoIcon from "@mui/icons-material/Tornado";

export const HeaderLogo = () => {
  return (
    <Typography variant="h3" sx={{ fontWeight: 'bold' }}>
      <TornadoIcon sx={{ fontSize: 38 }} />
      Squall
    </Typography>
  );
};
