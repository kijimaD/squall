import { createSlice } from "@reduxjs/toolkit";

export type View = {
  viewId: number,
}

export const viewSlice = createSlice({
  name: "view",
  initialState: {
    views: [] as View[],
  },
  reducers: {
    // id
    add: (state, action) => {
      const id = action.payload[0];
      const v: View = {viewId: id};
      state.views.push(v);
    },
  },
});

export const { add } = viewSlice.actions;

export default viewSlice.reducer;
