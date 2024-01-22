import { configureStore } from "@reduxjs/toolkit";
import counterReducer from "./counterSlice";
import viewReducer from "./viewSlice";

export const store = configureStore({
  reducer: {
    counter: counterReducer,
    view: viewReducer,
  },
});
4;
