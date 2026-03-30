let ws = null;
let localStream = null;
let peerConnection = null;
let currentTargetId = "";

const $ = (id) => document.getElementById(id);

const myIdInput = $("myId");
const targetIdInput = $("targetId");
const connectBtn = $("connectBtn");
const startCameraBtn = $("startCameraBtn");
const callBtn = $("callBtn");
const hangupBtn = $("hangupBtn");
const localVideo = $("localVideo");
const remoteVideo = $("remoteVideo");
const logBox = $("log");

// log("Connection success!") => <div>Connection success!</div>
function log(message) {
  logBox.textContent = message + "\n";
}
function sendSignal(toUserId, type, data) {
  if (!ws || ws.readyState !== WebSocket.OPEN) {
    return;
  }

  const payload = {
    to_user_id: toUserId,
    type: type,
    data: data,
  };

  ws.send(JSON.stringify(payload));
}
function createPeerConnection() {
  const pc = new RTCPeerConnection({
    iceServers: [{ urls: "stun:stun.l.google.com:19302" }],
  });

  pc.onicecandidate = function (event) {
    if (event.candidate && currentTargetId) {
      sendSignal(currentTargetId, "ice", {
        candidate: event.candidate,
      });
      log("Đã gửi ICE candidate");
    }
  };

  pc.ontrack = function (event) {
    log("Đã nhận remote stream");
    remoteVideo.srcObject = event.streams[0];
  };

  pc.onconnectionstatechange = function () {
    log("Connection state: " + pc.connectionState);
  };

  if (localStream) {
    localStream.getTracks().forEach(function (track) {
      pc.addTrack(track, localStream);
    });
  }

  return pc;
}
async function startCamera() {
  try {
    if (!navigator.mediaDevices || !navigator.mediaDevices.getUserMedia) {
      log(
        `Trình duyệt hiện tại không hỗ trợ camera trong ngữ cảnh này.
            Hãy dùng HTTPS hoặc localhost.`,
      );
      return;
    }

    localStream = await navigator.mediaDevices.getUserMedia({
      audio: true,
      video: true,
    });

    localVideo.srcObject = localStream;
    log("Đã mở camera/micro");
  } catch (error) {
    log("Lỗi mở camera: " + error.message);
  }
}
function addLocalTracksToPeerConnection() {
  if (!localStream) return;
  const senders = null;
  const existingTrackIds = null;
  localStream.getTracks().forEach();
}
function connectSignaling() {
  const myId = myIdInput.value.trim();
  if (!myId) {
    log("Vui lòng nhập ID của User 01");
    return;
  }
  const wsUrl =
    "ws://localhost:9999/ws/signaling?user=" + encodeURIComponent(myId);
  ws = new WebSocket(wsUrl);
  ws.onopen = function () {};
  ws.onmessage = async function (event) {
    log("Nhận signal: " + event.data);
    const msg = null;
    const fromUser = null;
    const type = null;
    const data = null;
    currentTargetId = fromUser;
    if (type === "offer") {
      log("Nhận offer từ " + fromUser);
      log("Đã gửi answer về " + fromUser);
    } else if (type === "answer") {
      log("Nhận answer từ " + fromUser);
    } else if (type === "ice") {
      log("Nhận ICE từ " + fromUser);
    }
  };
  ws.onerror = function (error) {
    log("WebSocket error");
    console.error(error);
  };
  ws.onclose = function () {
    log("WebSocket đã đóng");
  };
}
async function makeCall() {
  const targetId = targetIdInput.value.trim();
  if (!targetId) {
    log("Hãy nhập User ID để gọi");
    return;
  }
  if (!localStream) {
    log("Hãy mở camera trước");
    return;
  }
  if (!ws || ws.readyState !== WebSocket.OPEN) {
    log("Hãy kết nối signaling trước");
    return;
  }
  log("Đã gửi offer tới " + selectedFriend.username);
}
function hangUp() {
  if (peerConnection) {
  }
  log("Đã ngắt cuộc gọi");
}
