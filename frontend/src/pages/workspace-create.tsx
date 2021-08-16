import React from "react";
import axios from "axios";
import { SubmitHandler, useForm } from "react-hook-form";
import toast from "react-hot-toast";
import { useHistory } from "react-router-dom";

type Inputs = {
  image: string,
  name: string,
};

const WorkspaceCreate = () => {
  const history = useHistory();
  const { register, handleSubmit } = useForm();
  const onSubmit: SubmitHandler<Inputs> = (data) => {
    const myPromise = axios.post('/api/workspaces', data).then(() => { history.push('/'); });
    toast.promise(myPromise, {
      loading: 'Creating workspace...',
      success: 'Success!',
      error: 'Error when creating workspace.',
    });
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <div className="mb-3">
        <label htmlFor="name" className="form-label">Workspace's name</label>
        <input id="name" className="form-control" {...register("name", { required: true })} />
      </div>

      <div className="mb-3">
        <label htmlFor="image" className="form-label">Docker image</label>
        <input id="image" className="form-control" defaultValue="codercom/code-server" {...register("image", { required: true })} />
      </div>

      <input type="submit" className="btn btn-primary" />
    </form>
  );
};

export default WorkspaceCreate;
