# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Model.QuotaCustomDomains do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :available_domains,
    :can_use,
    :used_domains
  ]

  @type t :: %__MODULE__{
    :available_domains => integer() | nil,
    :can_use => boolean() | nil,
    :used_domains => integer() | nil
  }
end

defimpl Poison.Decoder, for: Ory.Model.QuotaCustomDomains do
  def decode(value, _options) do
    value
  end
end

