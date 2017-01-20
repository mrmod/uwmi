import React from 'react';
import {render} from 'react-dom';
import axios from 'axios';

class App extends React.Component {

  render () {
    console.log("APP");
    const projects = <Projects />;
    return <div>{projects}</div>
  }
}

class Project extends React.Component {
  render() {
    return(<div>
      <h1>{this.props.name}</h1>
      {this.props.description}
    </div>);
  }
}

class Projects extends React.Component {
  componentDidMount () {
    axios.get("/api/projects").then((response) => {
      console.log("FETCHED");
      this.setState({projects: []})
    });
  }
  render() {
    const projects = [];
    // const projects = this.state.projects.map((project) =>
    //   <Project name={project.name} description={project.description} />
    // );
    return <div><h1>Projects</h1>{projects}</div>
  }
}

render(<Projects/>, document.getElementById('app'));
