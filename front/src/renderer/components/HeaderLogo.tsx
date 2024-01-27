import { Typography } from "@mui/material";
import BubbleChartIcon from "@mui/icons-material/BubbleChart";

export const HeaderLogo = () => {
  return (
    <Typography variant="h3" sx={{ fontWeight: "bold" }}>
      <BubbleChartIcon sx={{ fontSize: 38 }} />
      Squall
    </Typography>
  );
};
