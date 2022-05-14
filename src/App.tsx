import React from 'react';
import {Themes} from "./distribute/position"
interface AppProps{
}
interface AppState{
}
function App(props: AppProps, state: AppState) {
  let theme = Themes.TextTheme
  return (
    <div className="App">
      <h1>Hello</h1>
    </div>
  );
}

export default App;
