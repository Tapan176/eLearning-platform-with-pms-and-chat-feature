import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Auth from './services/authService';
import Dashboard from './services/api';
import eLearning from './services/eLearningService';
import ProjectManagement from './services/projectManagementService';
import Chat from './services/chatService';

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
