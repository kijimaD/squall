import { IpcRenderer } from "electron";

declare global {
  interface Window {
    ipcRenderer: any;
    myAPI: any;
  }
}

export const { ipcRenderer } = window;
