import { createSlice } from "@reduxjs/toolkit";

export type View = {
  viewId: number;
  title: string;
  dataId: number;
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
    updateTitle: (state, action) => {
      state.views.forEach((v, i) => {
        if (action.payload.viewId == v.viewId) {
          state.views[i].title = action.payload.title;
        }
      });
    },
    remove: (state, action) => {
      const newViews = state.views.filter((n) => n.viewId !== action.payload);
      state.views = newViews;
    },
  },
});

export const { add, updateTitle, remove } = viewSlice.actions;

export default viewSlice.reducer;

// const v: View = { viewId: id };
