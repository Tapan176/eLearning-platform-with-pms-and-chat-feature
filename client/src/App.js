import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Auth from './components/Auth';
import Dashboard from './components/Dashboard';
import eLearning from './components/eLearning';
import ProjectManagement from './components/ProjectManagement';
import Chat from './components/Chat';

function App() {
  return (
    <Router>
      <Switch>
        <Route path="/auth" component={Auth} />
        <Route path="/dashboard" component={Dashboard} />
        <Route path="/elearning" component={eLearning} />
        <Route path="/projectmanagement" component={ProjectManagement} />
        <Route path="/chat" component={Chat} />
        {/* Add more routes as needed */}
      </Switch>
    </Router>
  );
}

export default App;
