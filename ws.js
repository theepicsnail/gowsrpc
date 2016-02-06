
r = Rpc()
r.onConnect = function(){
  r.call("Adder.Add", {
    "Num": 1
  }).then(function(reply) {
    // Add has completed, now let's request the new total
    return r.call("Adder.GetTotal")
  }).then(function(reply) {
    document.body.innerText = reply.Total + " views"
  })
}
