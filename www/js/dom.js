const setContentById = (id, content) => {
    const document = window.document.getElementById(id);
    if (document) {
        document.innerHTML = content;
    }
};

const formatTimezone = (tz) => {
    return tz.substring(tz.lastIndexOf("/") + 1)
} 

const setLocalTime = (time) => {
    setContentById('local-time', time);
};

const setLocalTimezone = (tz) => {
    setContentById('local-tz', formatTimezone(tz));
};

const setRemoteTime = (time) => {
    setContentById('remote-time', time);
};

const setRemoteTimezone = (tz) => {
    setContentById('remote-tz', formatTimezone(tz));
};

export default {
    setLocalTimezone,
    setLocalTime,
    setRemoteTime,
    setRemoteTimezone
};