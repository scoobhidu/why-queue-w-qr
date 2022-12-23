export const ActionCreator = (type, data) => {
  return {
    type: type,
    payload: data,
  };
};

export const newAction = (type) => {
  return {
    type: type,
  };
};
