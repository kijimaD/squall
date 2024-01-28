import { configureStore } from "@reduxjs/toolkit";
import viewReducer from "./viewSlice";
import {
  useSelector as rawUseSelector,
  TypedUseSelectorHook,
} from "react-redux";

export const store = configureStore({
  reducer: {
    view: viewReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;

// コンパイルエラーを回避するため、型情報付きのフックを作成する
export const myUseSelector: TypedUseSelectorHook<RootState> = rawUseSelector;
