import React from "react";
import { IWorkspace } from "../types/workspace";
import { BsArrowClockwise, BsStop, BsPlay, BsTrash, BsTerminal } from 'react-icons/bs';
import axios from "axios";
import toast from "react-hot-toast";
import { Link } from "react-router-dom";

import styles from './workspace-item.module.css';

type WorkspaceItemProps = {
  item: IWorkspace;
  refetch: () => Promise<any>;
}

const WorkspaceItem = ({ item, refetch }: WorkspaceItemProps) => {

  const color = (status: string) => {
    if (status === 'PENDING') {
      return { backgroundColor: 'rgb(245, 158, 11)' };
    } else if (status === 'running') {
      return { backgroundColor: 'rgb(16, 185, 129)' };
    } else if (status === 'exited') {
      return { backgroundColor: 'rgb(239, 68, 68)' };
    }
    return { backgroundColor: 'rgb(209, 213, 219)' };
  }
  
  const action = (status: string) => {
    if (status === 'PENDING') {
      return;
    } else if (status === 'running') {
      return (
        <>
          <a href="#" className="text-secondary fs-3 mx-2" onClick={stopWorkspace}>
            <BsStop />
          </a>
          <a href="#" className="text-secondary fs-3 mx-2" onClick={restartWorkspace}>
            <BsArrowClockwise />
          </a>
        </>
      );
    } else if (status === 'exited') {
      return (
        <a href="#" className="text-secondary fs-3 mx-2" onClick={startWorkspace}>
          <BsPlay />
        </a>
      );
    }
  }
  
  const deleteWorkspace = () => {
    const myPromise = axios.delete(`/api/workspaces/${item.id}`).then(refetch);
    toast.promise(myPromise, {
      loading: 'Deleting workspace...',
      success: 'Success!',
      error: 'Error when deleting workspace.',
    });
  }
  const startWorkspace = () => {
    const myPromise = axios.get(`/api/workspaces/${item.id}/start`).then(refetch);
    toast.promise(myPromise, {
      loading: 'Starting workspace...',
      success: 'Success!',
      error: 'Error when starting workspace.',
    });
  }
  const stopWorkspace = () => {
    const myPromise = axios.get(`/api/workspaces/${item.id}/stop`).then(refetch);
    toast.promise(myPromise, {
      loading: 'Stopping workspace...',
      success: 'Success!',
      error: 'Error when stopping workspace.',
    });
  }
  const restartWorkspace = () => {
    const myPromise = axios.get(`/api/workspaces/${item.id}/restart`).then(refetch);
    toast.promise(myPromise, {
      loading: 'Restarting workspace...',
      success: 'Success!',
      error: 'Error when restarting workspace.',
    });
  }
  
  return (
    <div className="shadow p-4 mb-3 rounded">
      <div className="d-flex align-items-center">
        <div className="rounded-circle me-3" style={{ ...color(item.status), height: '0.75rem', width: '0.75rem' }}></div>
        <div>
          { item.name }<br/>
          <span className="text-secondary">{ item.image }</span>
        </div>
        <div className="actions ms-auto">
          {action(item.status)}
          {item.status === 'running' && <Link to={`/workspace/${item.id}/terminal`} className="text-secondary fs-3 mx-2">
            <BsTerminal />
          </Link>}
          <a href="#" className="text-secondary fs-3 mx-2" onClick={deleteWorkspace}>
            <BsTrash />
          </a>
        </div>
      </div>
      {item.proxies.length !== 0 && item.status === "running" && (
        <>
          <hr />
          <div className="d-flex">
            {item.proxies.map((proxy) => (
              <a href={window.location.host === "env.maelgui.fr" ? `//${proxy.port}-${item.id}.env.maelgui.fr` : `/proxy/${item.id}/${proxy.port}/`} className={styles.proxy}>
                <div className={styles.proxyPort}>:{proxy.port}</div>
                <div className={styles.proxyName}>{proxy.name}</div>
              </a>
            ))}
          </div>
        </>
      )}
    </div>
  );
};

export default WorkspaceItem;
