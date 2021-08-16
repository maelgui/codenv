import React from 'react';
import axios from 'axios';
import { IWorkspace } from '../types/workspace';
import WorkspaceItem from '../components/workspace-item';
import { Link } from 'react-router-dom';



const WorkspaceList = () => {
  const [workspaces, setWorkspaces] = React.useState<Array<IWorkspace>>();

  React.useEffect(() => {
    axios.get('/api/workspaces').then((response) => {
      setWorkspaces(response.data);
    })
  }, []);

  return (
    <>
      <Link to="/workspace/add" className="float-end btn btn-success">Create a workspace</Link>
      <h1 className="mb-5">Workspaces</h1>
      <div>
        {workspaces?.map((w) => <WorkspaceItem item={w} />)}
      </div>
    </>
  )
}

export default WorkspaceList;
