"use client";

import { useEffect } from "react";
import {
  Cartesian3,
  createOsmBuildingsAsync,
  Ion,
  Math as CesiumMath,
  Terrain,
  Viewer,
} from "cesium";
import "cesium/Build/Cesium/Widgets/widgets.css";

export default function RealtimeVehicle() {
  useEffect(() => {
    // 環境変数からトークンを取得
    const cesiumToken = process.env.NEXT_PUBLIC_CESIUM_ION_ACCESS_TOKEN;
    console.log("Env: NEXT_PUBLIC_TEST= ", process.env.NEXT_PUBLIC_TEST);
    if (!cesiumToken) {
      console.error("Cesium Ion Access Token is missing.");
      return;
    }
    Ion.defaultAccessToken = cesiumToken;

    // ✅ `CESIUM_BASE_URL` を Next.js の `public/Cesium/` に設定
    (window as any).CESIUM_BASE_URL = "/Cesium/";

    async function initCesium() {
      const viewer = new Viewer("cesiumContainer", {
        terrain: await Terrain.fromWorldTerrain(),
      });

      viewer.camera.flyTo({
        destination: Cartesian3.fromDegrees(-122.4175, 37.655, 400),
        orientation: {
          heading: CesiumMath.toRadians(0.0),
          pitch: CesiumMath.toRadians(-15.0),
        },
      });

      const buildingTileset = await createOsmBuildingsAsync();
      viewer.scene.primitives.add(buildingTileset);
    }

    initCesium();

    return () => {
      const container = document.getElementById("cesiumContainer");
      if (container) {
        container.innerHTML = "";
      }
    };
  }, []);

  return (
    <div id="cesiumContainer" style={{ width: "100%", height: "100vh" }} />
  );
}
