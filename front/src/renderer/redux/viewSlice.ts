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
    add: (state, action) => {
      state.views.push(action.payload);
    },
    update: (state, action) => {
      state.views.forEach((v, i) => {
        if (action.payload.viewId == v.viewId) {
          state.views[i] = action.payload;
        }
      });
    },
  },
});

export const { add, update } = viewSlice.actions;

export default viewSlice.reducer;
