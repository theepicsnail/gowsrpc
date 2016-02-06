function Rpc() {
  if(this == window) {
    console.log("Please use 'new' in call to Rpc");
    return new Rpc()
  }
  console.log(this);
  this._sock = new WebSocket("ws://" + window.location.host + "/rpc");
  this._sock.onopen = this._onopen.bind(this)
  this._sock.onmessage = this._onmessage.bind(this)
  this._nid = 1;
  this._responseHandlers = {};
}

Rpc.prototype._onopen = function(evt) {
  if(this.onConnect)
    this.onConnect()
}

Rpc.prototype._onmessage = function(evt) {
  console.log(evt)

  response = JSON.parse(evt.data)
  handlers = this._responseHandlers[response.id]
  if(response.error) {
    handlers.reject(response.error)
  } else {
    handlers.accept(response.result)
  }
}

Rpc.prototype._nextId = function(evt) {
  this._nid += 1
  return this._nid
}

Rpc.prototype.call = function(method, args) {
  if(args == undefined)
    args = {}

  requestId = this._nextId()
  rpc = this
  return new Promise(function(accept, reject) {
    rpc._responseHandlers[requestId] = {
      "reject": reject, "accept": accept}
    rpc._sock.send(JSON.stringify({
      "method": method,
      "params": [args],
      "id": requestId
    }))
  });
}



