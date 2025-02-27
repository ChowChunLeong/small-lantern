"use client";

import React from "react";

type ImageProps = {
  imageUrl: string;
  alt: string;
};

export default function ProfileImage({ imageUrl, alt }: ImageProps) {
  return <img src={imageUrl} alt={alt} loading="lazy" />;
}
