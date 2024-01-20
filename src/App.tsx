export const App = () => {
    return (
        <div className="container">
            <h1>squall</h1>
            <button onClick={() => {
                window.myAPI.invoke("tab1");
            }}>google</button>
            <button onClick={() => {
                window.myAPI.invoke("tab2");
            }}>amazon</button>
        </div>
    );
};
