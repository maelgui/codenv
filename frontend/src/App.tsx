import React from 'react';
import {
  HashRouter as Router,
  Switch,
  Route,
} from 'react-router-dom';
import Layout from './components/layout';
import WorkspaceList from './pages/workspace-list';
import Terminal from './pages/terminal'

import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import WorkspaceCreate from './pages/workspace-create';
import { Toaster } from 'react-hot-toast';

function App() {
  return (
    <Router>
      <Toaster />
      <Layout>
        <Switch>
          <Route exact path="/">
            <WorkspaceList />
          </Route>
          <Route path="/workspace/add">
            <WorkspaceCreate />
          </Route>
          <Route path="/workspace/:id/terminal">
            <Terminal />
          </Route>
        </Switch>
      </Layout>
    </Router>
  );
}

export default App;
