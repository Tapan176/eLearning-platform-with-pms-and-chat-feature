import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Auth from './Auth/Auth';
import Dashboard from './Dashboard/Dashboard';
import eLearning from './eLearning/eLearning';
import ProjectManagement from './ProjectManagement/ProjectManagement';
import Chat from './Chat/Chat';

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
