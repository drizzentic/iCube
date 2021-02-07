import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import CharacterTable from './components/layout/Table/table.character';
import BooksTable from './components/layout/Table/table.books';
import HousesTable from './components/layout/Table/table.houses';
import FetchDataContextProvider from './components/context/fetchDataContext';


ReactDOM.render(
  <React.StrictMode>
    <FetchDataContextProvider >
      <Router>
        <Switch>
          <Route exact path="/" component={App}></Route>
          <Route path="/characters/" component={CharacterTable} />
          <Route path="/books" component={BooksTable} />
          <Route path="/houses" component={HousesTable} />
          <Route path="*">Not found</Route>
        </Switch>
      </Router>
    </FetchDataContextProvider>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
