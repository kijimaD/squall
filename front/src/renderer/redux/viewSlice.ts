import { createSlice } from "@reduxjs/toolkit";

export const viewSlice = createSlice({
  name: "view",
  initialState: {
    ids: [] as number[],
  },
  reducers: {
    // id
    add: (state, action) => {
      const id = action.payload[0];
      state.ids.push(id);
    },
  },
});

export const { add } = viewSlice.actions;

export default viewSlice.reducer;
