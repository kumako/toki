const fetchTimeZone = async () => {
    const res = await fetch("/api/timezone");
    const json = await res.json();
    return json["timezone"];
};

export default {
    fetchTimeZone
};
