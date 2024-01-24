import {
  app,
  BrowserWindow,
  BrowserView,
  ipcMain,
  webContents,
} from "electron";
import path from "path";
// This allows TypeScript to pick up the magic constants that's auto-generated by Forge's Webpack
// plugin that tells the Electron app where to look for the Webpack-bundled app code (depending on
// whether you're running in development or production).
declare const MAIN_WINDOW_WEBPACK_ENTRY: string;
declare const MAIN_WINDOW_PRELOAD_WEBPACK_ENTRY: string;

// Handle creating/removing shortcuts on Windows when installing/uninstalling.
if (require("electron-squirrel-startup")) {
  app.quit();
}

let mainWindow: any;
// UIを表示しているview
let uiView: any;
// 最も上にあって表示中のview
// viewをウィンドウに追加したときにそのviewに表示が切り替わってしまうので、記録しておいて戻すのに使う
let topView: any;

const createWindow = (): void => {
  mainWindow = new BrowserWindow({
    width: 1980,
    height: 1080,
    webPreferences: {
      // タブ領域の renderer プロセスから main プロセスにメッセージを送信するための contextBridge
      // preload: MAIN_WINDOW_PRELOAD_WEBPACK_ENTRY,
      preload: path.join(__dirname, "../renderer/main_window/preload.js"),
    },
    frame: false,
  });
  mainWindow.setBackgroundColor("gray");
  {
    const view = new BrowserView({
      webPreferences: {
        preload: path.join(__dirname, "../renderer/main_window/preload.js"),
      },
    });
    mainWindow.addBrowserView(view);
    view.webContents.loadURL(MAIN_WINDOW_WEBPACK_ENTRY);
    view.setBackgroundColor("white");
    const bound = mainWindow.getBounds();
    view.setBounds({ x: 0, y: 0, width: 1080, height: bound.height });
    view.webContents.openDevTools();
    uiView = view;
    topView = view;
  }
};

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.on("ready", createWindow);

// Quit when all windows are closed, except on macOS. There, it's common
// for applications and their menu bar to stay active until the user quits
// explicitly with Cmd + Q.
app.on("window-all-closed", () => {
  if (process.platform !== "darwin") {
    app.quit();
  }
});

app.on("activate", () => {
  // On OS X it's common to re-create a window in the app when the
  // dock icon is clicked and there are no other windows open.
  if (BrowserWindow.getAllWindows().length === 0) {
    createWindow();
  }
});

// IPCメインプロセスでのリクエストを待機
ipcMain.handle("openNewView", (e, arg) => {
  // debug
  const id = createView(arg.url);
  return id;
});

ipcMain.handle("changeTab", (e, arg) => {
  switchView(arg.id);
});

ipcMain.handle("changeHome", (e, arg) => {
  switchView(uiView.webContents.id);
});

ipcMain.handle("getTitleById", (e, arg) => {
  const title = getViewTitle(arg.id);
  return title;
});

function createView(url: string): number {
  const view = new BrowserView({
    webPreferences: {
      preload: path.join(__dirname, "../renderer/main_window/preload.js"),
    },
  });
  mainWindow.addBrowserView(view);
  mainWindow.setTopBrowserView(topView); // 新しく追加されたviewを表示しないように切り替える
  view.webContents.loadURL(url);
  const bound = mainWindow.getBounds();
  view.setBounds({ x: 300, y: 0, width: bound.width, height: bound.height });
  view.webContents.on("dom-ready", () => {
    // ページの読み込みが完了したらRendererプロセスにメッセージを送信
    uiView.webContents.send("pageLoaded", [
      view.webContents.id,
      view.webContents.getTitle(),
    ]);
  });
  return view.webContents.id;
}

function switchView(id: number) {
  const views = mainWindow
    .getBrowserViews()
    .filter((view) => view.webContents.id == id);
  console.assert(views.length === 1);
  mainWindow.setTopBrowserView(views[0]);
  topView = views[0];
}

function getViewTitle(id: number): string {
  const target = webContents.fromId(id);
  return target.getTitle();
}
