import { Grid, Button } from "@mui/material";

type Props = {
  title: string;
  url: string;
};

export const EntryButton = (props: Props) => {
  const { url } = props;
  const { title } = props;

  return (
    <Grid item xs={12} sm={6} spacing={1}>
      <Grid container direction="column">
        <Button
          color="black"
          style={{ justifyContent: "flex-start" }}
          onClick={() => {
            window.myAPI.invoke("changeTab", { url: url });
          }}
        >
          {title}
        </Button>
      </Grid>
    </Grid>
  );
};
