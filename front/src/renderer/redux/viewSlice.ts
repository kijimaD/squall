import { createSlice } from "@reduxjs/toolkit";

export type View = {
  viewId: number;
  title: string;
};

export const viewSlice = createSlice({
  name: "view",
  initialState: {
    views: [] as View[],
  },
  reducers: {
    // id
    add: (state, action) => {
      const id: number = action.payload[0];
      const title: string = action.payload[1];
      const v: View = { viewId: id, title: title };
      state.views.push(v);
    },
  },
});

export const { add } = viewSlice.actions;

export default viewSlice.reducer;
