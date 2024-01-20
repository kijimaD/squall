// See the Electron documentation for details on how to use preload scripts:
// https://www.electronjs.org/docs/latest/tutorial/process-model#preload-scripts

import { contextBridge, ipcRenderer } from "electron";
process.once("loaded", () => {
  contextBridge.exposeInMainWorld("myAPI", {
    ...ipcRenderer,
    on: (channel: string, func: (arg: any) => void) => {
      ipcRenderer.on(channel, (event, arg) => func(arg));
    },
  });
});
