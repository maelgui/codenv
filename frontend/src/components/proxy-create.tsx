import React from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { IProxy } from "../types/proxy";

type ProxyCreateProps = {
    onSubmit: SubmitHandler<IProxy>;
};

const ProxyCreate = ({ onSubmit }: ProxyCreateProps) => {
  const { register, handleSubmit } = useForm();

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <div className="mb-3">
        <label htmlFor="name" className="form-label">Proxy's name</label>
        <input id="name" className="form-control" {...register("name", { required: true })} />
      </div>

      <div className="mb-3">
        <label htmlFor="image" className="form-label">Port</label>
        <input id="image" className="form-control" type="number" {...register("port", { required: true, valueAsNumber: true, })} />
      </div>

      <input type="submit" className="btn btn-primary" />
    </form>
  );
};

export default ProxyCreate;
